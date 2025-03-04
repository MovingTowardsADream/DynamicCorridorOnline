import * as React from "react";
import style from './Navbar.module.css'
import {useNavigate} from "react-router-dom";
import {useContext} from "react";
import {AuthContext} from "../../context/AuthContext.tsx";

function Navbar() {
    const navigate = useNavigate();
    // eslint-disable-next-line @typescript-eslint/ban-ts-comment
    // @ts-expect-error
    const [auth, setAuth] = useContext(AuthContext);

    const handleStart = () => {
        navigate('/')
    }

    const handleSignUp = (e: React.MouseEvent<HTMLButtonElement>) => {
        e.preventDefault();
        navigate('/sign-up')
    }

    const handleLogIn = (e: React.MouseEvent<HTMLButtonElement>) => {
        e.preventDefault();
        navigate('/login')
    }

    const handleLogOut = (e: React.MouseEvent<HTMLButtonElement>) => {
        e.preventDefault();
        setAuth(null)
        localStorage.clear()
        navigate('/')
    }

    return (
        <>
            <nav className={style.navbar}>
                <div onClick={handleStart} style={{ cursor: 'pointer' }}>
                    <h1 className={style.logoText}>Dynamic Corridor</h1>
                </div>
                {auth ?
                    <div>
                        <button className={style.button} onClick={handleLogOut}>Log Out</button>
                    </div> :
                    <div>
                        <button className={style.button} onClick={handleSignUp}>Sign Up</button>
                        <button className={style.button} onClick={handleLogIn}>Log In</button>
                    </div>
                }
            </nav>
        </>
    );
}

export default Navbar;