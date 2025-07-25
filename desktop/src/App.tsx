import {RefObject} from "react";
import "./App.css";
import {useXTerm} from "react-xtermjs";

const MyTerminal = () => {
    const {instance, ref} = useXTerm()
    instance?.writeln('Hello from react-xtermjs!')
    instance?.onData((data) => instance?.write(data))

    return <div ref={ref as RefObject<HTMLDivElement>} style={{width: '100%', height: '100%'}}/>
}

function App() {
    // const [greetMsg, setGreetMsg] = useState("");
    // const [name, setName] = useState("");

    // async function greet() {
    //     Learn more about Tauri commands at https://tauri.app/develop/calling-rust/
        // setGreetMsg(await invoke("greet", {name}));
    // }

    return (
        <main className="container">
            <MyTerminal/>
        </main>
    );
}

export default App;
