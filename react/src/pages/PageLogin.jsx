import React, { useEffect } from 'react'
import { Card, Button, Label, TextInput } from 'flowbite-react';
import { useRecoilState } from "recoil";
import { loadingState, pagetTitleState } from '../state';
import ls from 'localstorage-slim';
import axios from 'axios';
import Swal from 'sweetalert2'
import { useForm } from "react-hook-form";
import { useNavigate } from "react-router-dom";



export const PageLogin = () => {
  const { register, formState: { errors }, handleSubmit, setValue: setValueForm } = useForm();
  const [pagetTitle, setPageTitle] = useRecoilState(pagetTitleState)
  const [loading, setLoading] = useRecoilState(loadingState)
  let navigate = useNavigate();


  useEffect(function () {
    setPageTitle("Login");
    setLoading(false);
  }, [])

  const onSubmitLogin = async (data) => {
    setLoading(true)
    axios.post('login', data).then(function (response) {
      if (response.status === true) {
        Swal.fire({
          title: "Berhasil",
          text: response.message,
          icon: "success"
        }).then(function () {
          ls.set(import.meta.env.VITE_APP_ID + '.access', response.data.access_token, { encrypt: false })
          const user = {
            email: response.data.email,
            uuid: response.data.uuid,
            nama_pengguna: response.data.nama_pengguna
          }
          ls.set(import.meta.env.VITE_APP_ID + '.user', user, { encrypt: false })
          navigate('/')
        });
      }else{
        Swal.fire({
          title: "Opps",
          text: response.message,
          icon: "warning"
        })
      }
       
      setLoading(false)
    })
  }

  return (
    <div className="flex flex-col items-center justify-center">
      <Card className='w-full sm:w-[200px] md:w-[400px] shadow-sm mt-8'>
        <form className="flex max-w-md flex-col gap-4">
          <h4 className='font-bold mb-3 text-center'>Silahkan Login Terlebih Dahulu</h4>
          <div>
            <div className="mb-2 block">
              <Label htmlFor="email" value="Masukan email" />
            </div>
            <TextInput
              {...register("email", {
                required: "Email wajid disi",
                pattern: {
                  value: /^[a-zA-Z0-9._-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$/,
                  message: 'email tidak valid',
                }
              })}
              id="email" type="email" placeholder="example@gmail.com" required />
            <ul>
              {errors.email ? <li className='TextError'>{errors.email.message} </li> : null}
            </ul>
          </div>
          <div>
            <div className="mb-2 block">
              <Label htmlFor="password" value="Masukan password" />
            </div>
            <TextInput
              {...register("password", {
                required: "Password wajib diisi",
              })}
              id="password" type="password" required />
            <ul>
              {errors.password ? <li className='TextError'>{errors.password.message} </li> : null}
            </ul>
          </div>
          <Button disabled={loading} onClick={handleSubmit(onSubmitLogin)} type="button">Login</Button>
        </form>
      </Card>
    </div>
  )
}
