import {useState, useContext, useEffect} from "react";
import {AuthContext} from "../../context/AuthContext.tsx";
import { FaCrown, FaMedal, FaAward } from 'react-icons/fa';
import style from "./LeaderBoard.module.css"

interface Statistic {
    id: string;
    username: string;
    expValue: number;
}

function LeaderBoard() {
    const [statistic, setStatistic] = useState<Statistic[]>([]);
    // eslint-disable-next-line @typescript-eslint/ban-ts-comment
    // @ts-expect-error
    // eslint-disable-next-line @typescript-eslint/no-unused-vars
    const [auth, setAuth] = useContext(AuthContext)

    const fetchBoard = async () => {
        const reqOptionsLeaderBoard = {
            method: "GET",
            headers: {
                'Authorization': `Bearer ${auth}`,
                'Content-Type': "application/json",
            },
            credentials: 'include' as RequestCredentials
        };

        const resp = await fetch("http://localhost:8080/api/v1/statistic/players?limit=10", reqOptionsLeaderBoard);
        const data = await resp.json();
        if (resp.status !== 200) {
            return
        }

        const dataLeaders = data["Leaders"]

        const statistics: Statistic[] = dataLeaders.map((item: any) => ({
            id: item["ID"],
            username: item["Username"],
            expValue: item["ExpValue"],
        }));

        setStatistic(statistics)
    }

    useEffect(() => {
        fetchBoard()
    }, []);

    return (
        <>
            <div>
                <h1 style={{ color: '#BFC0C0' }}>Leaderboard</h1>
                <table className={style.table}>
                    <thead>
                    <tr className={style.header}>
                        <th className={style.cell}>Rank</th>
                        <th className={style.cell}>Username</th>
                        <th className={style.cell}>Experience</th>
                    </tr>
                    </thead>
                    <tbody>
                    {statistic.map((stat: Statistic, index: number) => (
                        <tr key={stat.id} className={style.row}>
                            <td className={style.cell} data-label="Rank">
                                {index === 0 ? <FaCrown color="gold" /> :
                                    index === 1 ? <FaMedal color="silver" /> :
                                        index === 2 ? <FaAward color="bronze" /> :
                                            index + 1}
                            </td>
                            <td className={style.cell} data-label="Username">{stat.username}</td>
                            <td className={style.cell} data-label="Experience">{stat.expValue}</td>
                        </tr>
                    ))}
                    </tbody>
                </table>
            </div>
        </>
    )
}

export default LeaderBoard;