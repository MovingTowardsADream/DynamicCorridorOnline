import {ChangeEvent, useState} from 'react';
import Navbar from "../../components/Navbar/Navbar.tsx";
import style from "./Login.module.css"
import * as React from "react";
import ModalWindow from "../../components/ModalWindow/ModalWindow.tsx";

const defaultError: string = "something went wrong"

function Login() {
    const [login, setLogin] = useState<string>('');
    const [password, setPassword] = useState<string>('');
    const [rememberMe, setRememberMe] = useState(false);
    const [error, setError] = useState("")
    const [isModalOpen, setIsModalOpen] = useState(false);

    const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        console.log("try login");
        submitLogin();
    }

    const submitLogin = async () => {
        const reqOptionsSignIn = {
            method: "POST",
            headers: {
                'Content-Type': "application/json",
            },
            body: JSON.stringify({
                username: login,
                password: password
            }),
            credentials: 'include' as RequestCredentials
        };

        const respSignIn = await fetch("http://localhost:8080/user/auth/sign-in", reqOptionsSignIn);
        const dataLogin = await respSignIn.json();
        if (respSignIn.status !== 200) {
            let errData: string = dataLogin["error"]
            if (!errData) {
                errData = defaultError
            }

            errLogic(errData)
            return
        }
        console.log(dataLogin)
    }

    const errLogic = (err: string) => {
        setError(err)
        console.log("bad request")
        console.log("error: ", error)
        openErrWindow()
    }

    const openErrWindow = () => {
        setIsModalOpen(true)
    }

    return (
        <>
            <Navbar/>
            <div className={style.container}>
                <h1 className={style.titleText}>Login to Dynamic Corridor</h1>
                <form onSubmit={handleSubmit}>
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
                    <div className={style.rememberMeContainer}>
                        <label>
                            <input
                                type="checkbox"
                                checked={rememberMe}
                                onChange={(e: ChangeEvent<HTMLInputElement>) => setRememberMe(e.target.checked)}
                            />
                            Remember me
                        </label>
                    </div>
                    <div>
                        <button type="submit" className={style.button}>Login</button>
                    </div>
                </form>
            </div>
            {error !== "" &&
                <ModalWindow
                isOpen={isModalOpen}
                onClose={() => {setIsModalOpen(false);}}
                message={error}
                />
            }
        </>
    )
}

export default Login;