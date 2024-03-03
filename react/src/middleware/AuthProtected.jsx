
import { Navigate, Outlet } from 'react-router-dom'
import ls from 'localstorage-slim';

export const AuthProtected = () => {
  const userStorage = ls.get(import.meta.env.VITE_APP_ID + '.user', { decrypt: false });
  if (userStorage) {
    return (
      userStorage.uuid ? <Outlet /> : <Navigate to='/login' />
    )
  }
  return <Navigate to='/login' />
}


