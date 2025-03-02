import Navbar from "../../components/Navbar/Navbar.tsx"
import style from "./Start.module.css"

function Start() {
    return (
        <>
            <Navbar></Navbar>
            <div>
                <h1 className={style.title}><span className={style.cubeSpan}>Online</span> Dynamic Corridor</h1>

                <p className={style.text}>
                    <strong>Dynamic Corridor Online</strong> — this is not just a game, it’s a whole universe where everyone can find something of their own. Whether you're an experienced player or a newbie, exciting adventures and new friends await you.
                </p>
            </div>
        </>
    );
}

export default Start;