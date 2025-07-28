import React from "react";
import ReactDOM from "react-dom/client";
import App from "./App";
import {HeroUIProvider} from "@heroui/react";
import { getCurrentWindow } from "@tauri-apps/api/window";
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
ReactDOM.createRoot(document.getElementById("root") as HTMLElement).render(
  <React.StrictMode>
    <HeroUIProvider>
      <App/>
    </HeroUIProvider>
  </React.StrictMode>,
);
