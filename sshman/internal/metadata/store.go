package metadata

import (
	"encoding/json"
	"errors"
	"go.etcd.io/bbolt"
	"sman/internal/model"
	"time"
)

type Store struct {
	db *bbolt.DB
}

const bucketName = "ssh_connections"

func NewStore(path string) (*Store, error) {
	db, err := bbolt.Open(path, 0600, &bbolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return nil, err
	}

	// 初始化 bucket
	err = db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		return err
	})
	if err != nil {
		db.Close()
		return nil, err
	}

	return &Store{db: db}, nil
}

func (s *Store) Close() error {
	return s.db.Close()
}

// 添加或更新连接
func (s *Store) Save(conn *model.SSHConnection) error {
	conn.UpdatedAt = time.Now()
	if conn.CreatedAt.IsZero() {
		conn.CreatedAt = time.Now()
	}
	data, err := json.Marshal(conn)
	if err != nil {
		return err
	}

	return s.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		return b.Put([]byte(conn.ID), data)
	})
}

// 根据 ID 删除
func (s *Store) Delete(id string) error {
	return s.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		return b.Delete([]byte(id))
	})
}

// 根据 ID 获取
func (s *Store) Get(id string) (*model.SSHConnection, error) {
	var conn model.SSHConnection
	err := s.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		data := b.Get([]byte(id))
		if data == nil {
			return errors.New("connection not found")
		}
		return json.Unmarshal(data, &conn)
	})
	if err != nil {
		return nil, err
	}
	return &conn, nil
}

// 获取全部连接列表
func (s *Store) List() ([]*model.SSHConnection, error) {
	var conns []*model.SSHConnection
	err := s.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		return b.ForEach(func(k, v []byte) error {
			var conn model.SSHConnection
			if err := json.Unmarshal(v, &conn); err != nil {
				return err
			}
			conns = append(conns, &conn)
			return nil
		})
	})
	if err != nil {
		return nil, err
	}
	return conns, nil
}
