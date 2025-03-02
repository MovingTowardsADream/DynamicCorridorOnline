import {ChangeEvent, useState} from 'react';
import Navbar from "../../components/Navbar/Navbar.tsx";
import style from "./Signup.module.css"
import * as React from "react";

function Signup() {
    const [login, setLogin] = useState<string>('');
    const [password, setPassword] = useState<string>('');

    const handleSignup = (e: React.MouseEvent<HTMLButtonElement>) => {
        e.preventDefault();
        console.log("try sign up")
    }

    return (
        <>
            <Navbar/>
            <div className={style.container}>
                <h1 className={style.titleText}>Sign up to Dynamic Corridor</h1>
                <div>
                    <div className={style.inputContainer}>
                        <input type="text" value={login}
                               onChange={(e: ChangeEvent<HTMLInputElement>) => setLogin(e.target.value)}
                               placeholder="login">
                        </input>
                    </div>
                    <div className={style.inputContainer}>
                        <input type="password" value={password}
                               onChange={(e: ChangeEvent<HTMLInputElement>) => setPassword(e.target.value)}
                               placeholder="password">
                        </input>
                    </div>
                    <div>
                        <button className={style.button} onClick={handleSignup}>Sign Up</button>
                    </div>
                </div>
            </div>
        </>
    )
}

export default Signup;