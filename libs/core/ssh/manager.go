package ssh

import (
	"errors"
	"fmt"
	"github.com/daiyuang/sshman/core/metadata"
	"github.com/daiyuang/sshman/core/model"
	"github.com/samber/lo"
	gossh "golang.org/x/crypto/ssh"
	"sync"
	"time"
)

type Manager struct {
	mu          sync.RWMutex
	connections map[string]*model.SSHConnection
	clients     map[string]*gossh.Client
	store       *metadata.Store
}

func NewManager(store *metadata.Store) *Manager {
	list, err := store.List()
	if err != nil {
		return nil
	}
	connections := make(map[string]*model.SSHConnection)
	lo.ForEach(list, func(conn *model.SSHConnection, index int) {
		connections[conn.ID] = conn
	})
	return &Manager{
		connections: connections,
		clients:     make(map[string]*gossh.Client),
		store:       store,
	}
}

// Add 新增连接配置（只存配置，不连接）
func (m *Manager) Add(conn *model.SSHConnection) error {
	if conn == nil || conn.ID == "" {
		return errors.New("invalid connection")
	}
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, exists := m.connections[conn.ID]; exists {
		return errors.New("connection ID already exists")
	}
	conn.CreatedAt = time.Now()
	conn.UpdatedAt = time.Now()
	m.connections[conn.ID] = conn
	err := m.store.Save(conn)
	if err != nil {
		return err
	}
	return nil
}

// Remove 删除配置并断开连接（如果有）
func (m *Manager) Remove(id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if client, ok := m.clients[id]; ok {
		client.Close()
		delete(m.clients, id)
	}
	if _, ok := m.connections[id]; !ok {
		return errors.New("connection not found")
	}
	delete(m.connections, id)
	return nil
}

// Connect 基于配置创建 ssh.Client 并存储
func (m *Manager) Connect(id string) (*gossh.Client, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	conn, ok := m.connections[id]
	if !ok {
		return nil, errors.New("connection not found")
	}
	if client, ok := m.clients[id]; ok {
		// 已有连接，返回
		return client, nil
	}

	sshConfig := &gossh.ClientConfig{
		User:            conn.Username,
		HostKeyCallback: gossh.InsecureIgnoreHostKey(), // TODO: 生产环境要验证 HostKey
		Timeout:         10 * time.Second,
	}

	// 认证方式
	switch conn.AuthMethod {
	case "password":
		sshConfig.Auth = []gossh.AuthMethod{
			gossh.Password(conn.Password),
		}
	case "key":
		signer, err := gossh.ParsePrivateKey(conn.PrivateKey)
		if err != nil {
			return nil, fmt.Errorf("failed to parse private key: %v", err)
		}
		sshConfig.Auth = []gossh.AuthMethod{
			gossh.PublicKeys(signer),
		}
	default:
		return nil, errors.New("unsupported auth method")
	}

	addr := fmt.Sprintf("%s:%d", conn.Host, conn.Port)
	client, err := gossh.Dial("tcp", addr, sshConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to dial: %v", err)
	}

	m.clients[id] = client
	return client, nil
}

// Close 关闭并移除连接
func (m *Manager) Close(id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	client, ok := m.clients[id]
	if !ok {
		return errors.New("connection not found or not connected")
	}
	err := client.Close()
	delete(m.clients, id)
	return err
}

// GetClient 获取活跃 ssh.Client
func (m *Manager) GetClient(id string) (*gossh.Client, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	c, ok := m.clients[id]
	return c, ok
}

// ListConnections 返回所有连接配置
func (m *Manager) ListConnections() []*model.SSHConnection {
	m.mu.RLock()
	defer m.mu.RUnlock()
	list := make([]*model.SSHConnection, 0, len(m.connections))
	for _, v := range m.connections {
		list = append(list, v)
	}
	return list
}
