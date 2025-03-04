import {useContext} from 'react';
import {AuthContext} from "../../context/AuthContext";
import Navbar from "../../components/Navbar/Navbar.tsx";
import LeaderBoard from "../../components/LeaderBoard/LeaderBoard.tsx";
import style from "./Lobby.module.css"

function Lobby() {
    // eslint-disable-next-line @typescript-eslint/ban-ts-comment
    // @ts-expect-error
    // eslint-disable-next-line @typescript-eslint/no-unused-vars
    const [auth, setAuth] = useContext(AuthContext)

    return (
        <>
            <Navbar />
            <div className={style.container}>
                <div className="d-flex">
                    <div>
                        <div>
                            <p>Username:</p>
                            <p>Experience:</p>
                            <button></button>
                        </div>
                    </div>
                    <div>
                        <LeaderBoard />
                    </div>
                </div>
            </div>
        </>
    )
}

export default Lobby;