import {Routes, Route} from "react-router-dom";
import './App.css'
import Login from "./pages/Login/Login.tsx";
import Signup from "./pages/Signup/Signup.tsx";
import Lobby from "./pages/Lobby/Lobby.tsx";

function App() {
  return (
    <>
        <Routes>
            <Route path="">
                <Route path='/sign-up' element={<Signup/>}/>
                <Route path='/login' element={<Login/>}/>
                <Route path='/' element={<Lobby/>}/>
            </Route>
        </Routes>
    </>
  )
}

export default App
