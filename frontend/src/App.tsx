import {Routes, Route} from "react-router-dom";
import './App.css'
import Start from "./pages/Start/Start.tsx";

function App() {
  return (
    <>
        <Routes>
            <Route path="">
                <Route path='/' element={<Start/>}/>
            </Route>
        </Routes>
    </>
  )
}

export default App
