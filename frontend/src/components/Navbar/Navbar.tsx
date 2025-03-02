import * as React from "react";
import style from './Navbar.module.css'
import {useNavigate} from "react-router-dom";

function Navbar() {
    const navigate = useNavigate();

    const handleStart = () => {
        console.log("navigate on /")
        navigate('/')
    }

    const handleSignUp = (e: React.MouseEvent<HTMLButtonElement>) => {
        e.preventDefault();
        console.log("navigate on /sign-up")
        navigate('/sign-up')
    }

    const handleLogIn = (e: React.MouseEvent<HTMLButtonElement>) => {
        e.preventDefault();
        console.log("navigate on /login")
        navigate('/login')
    }

    return (
        <>
            <nav className={style.navbar}>
                <div onClick={handleStart} style={{ cursor: 'pointer' }}>
                    <h1 className={style.logoText}>Dynamic Corridor</h1>
                </div>
                <div>
                    <button className={style.button} onClick={handleSignUp}>Sign Up</button>
                    <button className={style.button} onClick={handleLogIn}>Log In</button>
                </div>
            </nav>
        </>
    );
};

export default Navbar;