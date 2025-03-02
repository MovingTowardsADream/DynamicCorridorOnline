import { createContext, useEffect, useState } from 'react';

export const UserContext = createContext({});

function UserProvider(props) {
    const [user, setUser] = useState(null)

    useEffect(()=> {
        const userData = localStorage.getItem("userData")
        if (userData){
            const foundedUser = JSON.parse(userData)
            setUser(foundedUser)
        }
    }, [])

    return (
        <UserContext.Provider value={[user, setUser]}>
            {props.children}
        </UserContext.Provider>
    )
}

export default UserProvider;
