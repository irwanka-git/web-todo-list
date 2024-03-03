import { Dropdown, Navbar, Avatar } from 'flowbite-react'
import React from 'react'
import ls from 'localstorage-slim';

import { useNavigate } from "react-router-dom";


export const Header = () => {
    const userStorage = ls.get(import.meta.env.VITE_APP_ID + '.user', { decrypt: false }); 

    const navigate = useNavigate();
    const submitLogout = ()=>{
        ls.set(import.meta.env.VITE_APP_ID + '.access', null, { encrypt: false })
        ls.set(import.meta.env.VITE_APP_ID + '.user', null, { encrypt: false })
        ls.clear()
        navigate('/login')
    } 

    return (
        <div className='flex-1 my-2'>
            <Navbar fluid rounded className='px-4 rounded-md shadow-sm py-3'>
                <Navbar.Brand href='/'>
                    <span className="whitespace-nowrap text-md font-semibold">
                        <span>Todo List</span>
                    </span>
                </Navbar.Brand>

                {userStorage ?
                    <div className="flex md:order-2">
                        <Dropdown
                            arrowIcon={false}
                            inline
                            label={
                                <Avatar alt="User settings" img={`https://ui-avatars.com/api/?name=${userStorage.nama_pengguna}`} rounded />
                            }
                        >
                            <Dropdown.Header>
                                <span className="block text-sm font-bold text-green-500">{userStorage.email}</span>
                                <span className="block truncate text-sm font-medium">{userStorage.nama_pengguna}</span>
                            </Dropdown.Header>
                            <Dropdown.Item onClick={()=>submitLogout()}>Logout</Dropdown.Item>
                        </Dropdown>
                    </div> : <></>}
            </Navbar>
        </div>
    )
}
