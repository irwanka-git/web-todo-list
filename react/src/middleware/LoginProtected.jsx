 
import { Navigate, Outlet } from 'react-router-dom' 
import ls from 'localstorage-slim';

export const LoginProtected = () => { 
    const userStorage =  ls.get(import.meta.env.VITE_APP_ID + '.user', { decrypt: false });
    if (userStorage) {
         return (
             userStorage.status ?  <Navigate to='/'/>  : <Outlet/>
          )
    }
    return <Outlet/>
}


