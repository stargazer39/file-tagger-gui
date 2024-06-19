import { useState } from "react";
import { Greet, GetFileList } from "../wailsjs/go/main/App";

function App() {
  const [list, setList] = useState([]);

  async function greet() {
    setList(await GetFileList());
  }

  return (
    <div>
      <button onClick={greet}>Load files</button>
      {list.map((f) => (
        <div key={f.Name}>
          <div>Name {f.Name}</div>
        </div>
      ))}
    </div>
  );
}

export default App;
