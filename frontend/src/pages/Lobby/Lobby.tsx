import {useContext, useEffect, useState} from 'react';
import {AuthContext} from "../../context/AuthContext";
import Navbar from "../../components/Navbar/Navbar.tsx";
import LeaderBoard from "../../components/LeaderBoard/LeaderBoard.tsx";
import style from "./Lobby.module.css"

function Lobby() {
    // eslint-disable-next-line @typescript-eslint/ban-ts-comment
    // @ts-expect-error
    // eslint-disable-next-line @typescript-eslint/no-unused-vars
    const [auth, setAuth] = useContext(AuthContext)

    const [username, setUsername] = useState(null)
    const [expValue, setExpValue] = useState(null)

    const fetchExp = async () => {
        const reqOptionsExp = {
            method: "GET",
            headers: {
                'Authorization': `Bearer ${auth}`,
                'Content-Type': "application/json",
            },
            credentials: 'include' as RequestCredentials
        };

        const resp = await fetch("http://localhost:8080/api/v1/statistic/players/experience", reqOptionsExp);
        const data = await resp.json();
        if (resp.status !== 200) {
            return
        }

        setUsername(data["Username"])
        setExpValue(data["ExpValue"])
    }

    useEffect(() => {
        fetchExp()
    }, []);

    return (
        <>
            <Navbar />
            {auth ?
                <div className={style.container}>
                    <div className="d-flex">
                        <div>
                            <div>
                                <p>{username}</p>
                                <p>Experience: {expValue}</p>
                                <button>Play</button>
                            </div>
                        </div>
                        <div>
                            <LeaderBoard />
                        </div>
                    </div>
                </div> :
                <div>
                    <h1 className={style.title}><span className={style.cubeSpan}>Online</span> Dynamic Corridor</h1>

                    <p className={style.text}>
                        <strong>Dynamic Corridor Online</strong> — this is not just a game, it’s a whole universe where everyone can find something of their own. Whether you're an experienced player or a newbie, exciting adventures and new friends await you.
                    </p>
                </div>
            }

        </>
    )
}

export default Lobby;