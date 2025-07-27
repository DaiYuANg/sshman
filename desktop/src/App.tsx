import {RefObject} from "react";
import "./App.css";
import {useXTerm} from "react-xtermjs";
import {
  Button, Card, CardBody,
  Modal,
  ModalBody,
  ModalContent,
  ModalFooter,
  ModalHeader,
  Tab,
  Tabs,
  useDisclosure
} from "@heroui/react";
import {Panel, PanelGroup, PanelResizeHandle} from "react-resizable-panels";

const MyTerminal = () => {
  const {instance, ref} = useXTerm()
  instance?.writeln('Hello from react-xtermjs!')
  instance?.onData((data) => instance?.write(data))

  return <div ref={ref as RefObject<HTMLDivElement>} style={{width: '100%', height: '100%'}}/>
}
import { getCurrentWindow } from '@tauri-apps/api/window';

// when using `"withGlobalTauri": true`, you may use
// const { getCurrentWindow } = window.__TAURI__.window;

const appWindow = getCurrentWindow();

document
  .getElementById('titlebar-minimize')
  ?.addEventListener('click', () => appWindow.minimize());
document
  .getElementById('titlebar-maximize')
  ?.addEventListener('click', () => appWindow.toggleMaximize());
document
  .getElementById('titlebar-close')
  ?.addEventListener('click', () => appWindow.close());
const App = () => {
  const {isOpen, onOpen, onOpenChange} = useDisclosure();
  return (
    <main className="container h-screen w-full">
      <PanelGroup direction="horizontal" className="h-full">
        {/* 左侧栏 */}
        <Panel defaultSize={25} minSize={15} maxSize={50}>
          <div data-tauri-drag-region className="h-full mt-2  p-4  border-gray-300">
            {/*<h2 className="text-lg font-bold mb-2">左侧栏</h2>*/}
            <Button onPress={onOpen}>test button</Button>
          </div>
        </Panel>

        {/* 拖拽手柄 */}
        <PanelResizeHandle className="w-1 bg-gray-300 hover:bg-gray-400 cursor-col-resize"/>

        {/* 右侧主区域 */}
        <Panel>
          <div className="h-full p-4 overflow-auto bg-white">
            <Tabs aria-label="Options">
              <Tab key="photos" title="Photos">
                <h2 className="text-lg font-bold mb-2">终端</h2>
                <MyTerminal/>
              </Tab>
              <Tab key="music" title="Music">
                <Card>
                  <CardBody>
                    Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex
                    ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse
                    cillum dolore eu fugiat nulla pariatur.
                  </CardBody>
                </Card>
              </Tab>
              <Tab key="videos" title="Videos">
                <Card>
                  <CardBody>
                    Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt
                    mollit anim id est laborum.
                  </CardBody>
                </Card>
              </Tab>
            </Tabs>

          </div>
        </Panel>
      </PanelGroup>
      <Modal isOpen={isOpen} onOpenChange={onOpenChange}>
        <ModalContent>
          {(onClose) => (
            <>
              <ModalHeader className="flex flex-col gap-1">Modal Title</ModalHeader>
              <ModalBody>
                <p>
                  Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam pulvinar risus non
                  risus hendrerit venenatis. Pellentesque sit amet hendrerit risus, sed porttitor
                  quam.
                </p>
                <p>
                  Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam pulvinar risus non
                  risus hendrerit venenatis. Pellentesque sit amet hendrerit risus, sed porttitor
                  quam.
                </p>
                <p>
                  Magna exercitation reprehenderit magna aute tempor cupidatat consequat elit dolor
                  adipisicing. Mollit dolor eiusmod sunt ex incididunt cillum quis. Velit duis sit
                  officia eiusmod Lorem aliqua enim laboris do dolor eiusmod. Et mollit incididunt
                  nisi consectetur esse laborum eiusmod pariatur proident Lorem eiusmod et. Culpa
                  deserunt nostrud ad veniam.
                </p>
              </ModalBody>
              <ModalFooter>
                <Button color="danger" variant="light" onPress={onClose}>
                  Close
                </Button>
                <Button color="primary" onPress={onClose}>
                  Action
                </Button>
              </ModalFooter>
            </>
          )}
        </ModalContent>
      </Modal>
    </main>
  );
};

export default App;
