import {ChangeEvent, useState} from 'react';
import Navbar from "../../components/Navbar/Navbar.tsx";
import style from "./Signup.module.css"
import * as React from "react";
import ModalWindow from "../../components/ModalWindow/ModalWindow.tsx";
import {useNavigate} from "react-router-dom";

const defaultError: string = "something went wrong"

function Signup() {
    const navigate = useNavigate();
    const [error, setError] = useState("");
    const [login, setLogin] = useState<string>('');
    const [password, setPassword] = useState<string>('');
    const [isModalOpen, setIsModalOpen] = useState(false);

    const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        submitSignup().then(() => console.log("error submit sign up"));
    }

    const submitSignup = async () => {
        const reqOptionsSignUp = {
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

        const respSignUp = await fetch("http://localhost:8080/user/auth/sign-up", reqOptionsSignUp);
        const dataSignUp = await respSignUp.json();

        if (respSignUp.status !== 200) {
            let errData: string = dataSignUp["error"]
            if (!errData) {
                errData = defaultError
            }
            errLogic(errData)
            return
        }

        navigate("/");
    }

    const errLogic = (err: string) => {
        setError(err)
        console.log("bad request - error: ", error)
        openErrWindow()
    }

    const openErrWindow = () => {
        setIsModalOpen(true)
    }

    return (
        <>
            <Navbar/>
            <div className={style.container}>
                <h1 className={style.titleText}>Sign up to Dynamic Corridor</h1>
                <form onSubmit={handleSubmit}>
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
                            <button type="submit" className={style.button} >Sign Up</button>
                        </div>
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

export default Signup;