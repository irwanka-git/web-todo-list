import { useState } from 'react'
import { BrowserRouter, Routes, Route } from 'react-router-dom'
import { AuthProtected } from './middleware/AuthProtected'
import { PageTask } from './pages/PageTask'
import { PageLogin } from './pages/PageLogin'
import { BaseLayout } from './layouts/BaseLayout'
import axios from 'axios'
import ls from 'localstorage-slim';
import { LoginProtected } from './middleware/LoginProtected'

axios.interceptors.response.use(response => {
  return response.data;
}, error => {
  if (error.response) {
    console.log(error.response)
    return error.response.data
  } else {
    console.log("ERROR RESPONSE REFUSED")
    return
  }
});

axios.interceptors.request.use(function (config) {
  config.baseURL = import.meta.env.VITE_BASE_URL_API
  const appID = import.meta.env.VITE_APP_ID
  const storageValue = ls.get(appID + '.access', { decrypt: true });
  if (storageValue) {
      config.headers.Authorization = `Bearer ${storageValue}`;
  }
  return config;
});

function App() {

  return (
    <>
      <BrowserRouter>
        <Routes>
          <Route element={<BaseLayout />}>
            <Route element={<AuthProtected />}>
              <Route path='/' element={<PageTask />} />
            </Route>
            <Route element={<LoginProtected />}>
              <Route path='/login' element={<PageLogin />} />
            </Route>
          </Route>
        </Routes>
      </BrowserRouter>
    </>
  )
}

export default App
