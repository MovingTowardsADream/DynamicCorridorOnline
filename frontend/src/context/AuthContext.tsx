import { createContext, useEffect, useState } from 'react';

export const AuthContext = createContext({});

// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-expect-error
function AuthProvider(props) {
    const [auth, setAuth] = useState(null)

    useEffect(()=> {
        const authData = sessionStorage.getItem("authToken")
        if (authData){
            const foundedAuth = JSON.parse(authData)
            setAuth(foundedAuth)
        }
    }, [])

    return (
        <AuthContext.Provider value={[auth, setAuth]}>
            {props.children}
        </AuthContext.Provider>
    )
}

export default AuthProvider;
