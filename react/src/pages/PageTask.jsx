import React, { useEffect, useState } from 'react'
import { useRecoilState } from "recoil";
import { loadingState, pagetTitleState } from '../state';
import ls from 'localstorage-slim';
import axios from 'axios';
import Swal from 'sweetalert2'
import { useForm } from "react-hook-form";
import { useNavigate } from "react-router-dom";
import { ListLoading } from '../components/ListLoading';
import { Button, Card, Modal, TextInput, Textarea, Label } from 'flowbite-react';
import { HiOutlineBookmark, HiOutlineExclamationCircle, HiOutlinePencil, HiOutlinePlus, HiOutlineTrash } from "react-icons/hi";
import { format as formatTime } from "date-fns";



export const PageTask = () => {

    const [pagetTitle, setPageTitle] = useRecoilState(pagetTitleState)
    const [loading, setLoading] = useRecoilState(loadingState)

    const [listTask, setListTask] = useState([])

    const [openModalCreate, setOpenModalCreate] = useState(false);
    const { register: regFormCreate, formState: { errors: errorFormCreate }, handleSubmit: handleSubmitCreate, setValue: setValueFormCreate, reset: resetFormCreate } = useForm({ mode: "onBlur" });

    const [openModalDetil, setOpenModalDetil] = useState(false);
    const [detilTask, setDetilTask] = useState({})

    const [openModalDelete, setOpenModalDelete] = useState(false);
    const [deleteTaskID, setDeleteTaskID] = useState(0)

    const [openModalEdit, setOpenModalEdit] = useState(false);
    const { register: regFormEdit, formState: { errors: errorFormEdit }, handleSubmit: handleSubmitUpdate, setValue: setValueFormEdit, reset: resetFormEdit } = useForm({ mode: "onBlur" });
    const [editTask, setEditTask] = useState({})

    useEffect(function () {
        setPageTitle("Task");
        setLoading(false);
        fetchTasks();
    }, [])

    const fetchTasks = async => {
        setLoading(true);
        axios.get('/list-task').then(function (response) {
            if (response.status == true) {
                setListTask(response.data)
            } else {
                setListTask([])
            }
            setTimeout(() => { setLoading(false) }, 500)
        }).catch(function () {
            setLoading(false);
        });
    }

    const showFormCreate = () => {
        resetFormCreate()
        setValueFormCreate('title', '')
        setValueFormCreate('description', '')
        setOpenModalCreate(true)
    }

    const submitCreateTask = async (data) => {
        setOpenModalCreate(false)
        setLoading(true)
        axios.post('/create-task', data).then(function (response) {
            if (response.status === true) {
                Swal.fire({
                    title: "Berhasil",
                    text: response.message,
                    icon: "success"
                }).then(function () {
                    fetchTasks()
                });
            } else {
                Swal.fire({
                    title: "Opps",
                    text: response.message,
                    icon: "warning"
                })
            }
            setLoading(false)
        })
    }

    const getDetilTask = (id_task) => {
        axios.get('/get-detil-task/' + id_task).then(function (response) {
            if (response.status == true) {
                setDetilTask(response.data)
                setOpenModalDetil(true)
            } else {
                setDetilTask(null)
            }
        }).catch(function () {
            alert("Terjadi Kesalahan!")
        });
    }

    const showFormDelete = (id_task) => {
        setDeleteTaskID(id_task)
        setOpenModalDelete(true)
    }
    const submitDeleteTask = async => {
        setOpenModalDelete(false)
        axios.delete('delete-task/' + deleteTaskID).then(function (response) {
            if (response.status === true) {
                Swal.fire({
                    title: "Berhasil",
                    text: response.message,
                    icon: "success"
                }).then(function () {
                    fetchTasks()
                });
            } else {
                Swal.fire({
                    title: "Opps",
                    text: response.message,
                    icon: "warning"
                })
            }
            setDeleteTaskID(0)
        })
    }

    const showFormEdit = (id_task) => {
        axios.get('/get-detil-task/' + id_task).then(function (response) {
            if (response.status == true) {
                setEditTask(response.data)
                resetFormEdit()
                setValueFormEdit("title", response.data.title)
                setValueFormEdit("description", response.data.description)
                setOpenModalEdit(true)
            } else {
                setEditTask(null)
            }
        }).catch(function () {
            alert("Terjadi Kesalahan!")
        });
    }

    const submitUpdateTask = async(data) =>{
        setOpenModalEdit(false)
        setLoading(true)
        axios.patch('/update-task/' + editTask.id_task, data).then(function (response) {
            if (response.status === true) {
                Swal.fire({
                    title: "Berhasil",
                    text: response.message,
                    icon: "success"
                }).then(function () {
                    fetchTasks()
                });
            } else {
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
        <div>
            <div className='flex-1 mb-6'>
                <Button disabled={loading} type="button" onClick={() => showFormCreate()}>
                    <HiOutlinePlus className="mr-2 h-4 w-4" />
                    Task Baru
                </Button>
            </div>

            <Card className="max-w-full">
                <div className="mb-2 flex items-center justify-between">
                    <h5 className="text-lg font-bold leading-none text-gray-900 dark:text-white">Riwayat Task</h5>
                </div>
                <div className="flow-root">
                    {loading === true ? <ListLoading /> :
                        <ul className="divide-y divide-gray-200">
                            {
                                listTask.length > 0 ?
                                    <>
                                        {
                                            listTask.map((task) => (
                                                <li className="py-2 sm:py-3" key={task.id_task}>
                                                    <div className="flex items-center space-x-4">
                                                        <div className="shrink-0">
                                                            <HiOutlineBookmark className='h-6 w-6' />
                                                        </div>
                                                        <a href='#' onClick={() => getDetilTask(task.id_task)} className="min-w-0 flex-1">
                                                            <p className="truncate text-sm font-medium text-gray-900 dark:text-white">{task.title}</p>
                                                            <p className="truncate text-sm text-gray-500 dark:text-gray-400">{task.description}</p>
                                                        </a>
                                                        <div className="inline-flex items-center">
                                                            <Button size="xs" color="light" onClick={()=>showFormEdit(task.id_task)} outline className='mr-1'>
                                                                <HiOutlinePencil size={14} />
                                                            </Button>
                                                            <Button size="xs" color="failure" onClick={()=>showFormDelete(task.id_task)}>
                                                                <HiOutlineTrash size={14} />
                                                            </Button>
                                                        </div>
                                                    </div>
                                                </li>
                                            ))

                                        }
                                    </>
                                    :
                                    <>
                                        Belum ada task
                                    </>
                            }
                        </ul>
                    }
                </div>
            </Card>

            <form key={'form-create'}>
                <Modal id='modal-create' show={openModalCreate} onClose={() => setOpenModalCreate(false)}>
                    <Modal.Header>Buat Task Baru</Modal.Header>
                    <Modal.Body>
                        <div className="space-y-3">
                            <div>
                                <div className="mb-2 block">
                                    <Label htmlFor="judul" value="Judul" />
                                </div>
                                <TextInput
                                    {...regFormCreate("title", {
                                        required: "Judul wajib diisi",
                                    })}
                                    id="title" placeholder="Masukan judul" required />
                                <ul>
                                    {errorFormCreate.title ? <li className='TextError'>{errorFormCreate.title.message} </li> : null}
                                </ul>
                            </div>
                            <div>
                                <div className="mb-2 block">
                                    <Label htmlFor="deskripsi" value="Deskripsi" />
                                </div>
                                <Textarea
                                    {...regFormCreate("description", {
                                        required: "Deskripsi wajib diisi",
                                    })}
                                    id="description" rows={4} placeholder="Masukan deskripsi" required />
                                <ul>
                                    {errorFormCreate.description ? <li className='TextError'>{errorFormCreate.description.message} </li> : null}
                                </ul>
                            </div>
                        </div>

                    </Modal.Body>
                    <Modal.Footer>
                        <Button onClick={handleSubmitCreate(submitCreateTask)}>Simpan</Button>
                        <Button color="gray" onClick={() => setOpenModalCreate(false)}>
                            Batal
                        </Button>
                    </Modal.Footer>
                </Modal>
            </form>

            <form key={'form-detil'}>
                <Modal id='modal-detil' show={openModalDetil} onClose={() => setOpenModalDetil(false)}>
                    <Modal.Header>Detil Task #{detilTask.id_task}</Modal.Header>
                    <Modal.Body>
                        <div className="space-y-3">
                            <div>
                                <div className="mb-2 block">
                                    <Label htmlFor="judul" value="Judul" />
                                </div>
                                <div className='p-2 text-sm rounded-md border border-neutral-300'>
                                    {detilTask.title}
                                </div>
                            </div>
                            <div>
                                <div className="mb-2 block">
                                    <Label htmlFor="deskripsi" value="Deskripsi" />
                                </div>
                                <div className='p-2 text-sm rounded-md border border-neutral-300'>
                                    {detilTask.description}
                                </div>
                            </div>
                            <div>
                                <div className="mb-2 block">
                                    <Label htmlFor="created_at" value="Created At" />
                                </div>
                                <div className='p-2 text-sm rounded-md border border-neutral-300'>
                                    {detilTask.created_at ? formatTime(detilTask.created_at, "d MMM yyyy HH:mm:ss") : ""} WIB
                                </div>
                            </div>
                            {detilTask.created_at != detilTask.updated_at ? <div>
                                <div className="mb-2 block">
                                    <Label htmlFor="updated_at" value="Updated At" />
                                </div>
                                <div className='p-2 text-sm rounded-md border border-neutral-300'>
                                    {detilTask.created_at ? formatTime(detilTask.updated_at, "d MMM yyyy HH:mm:ss") : ""} WIB
                                </div>
                            </div> : <></>}

                        </div>

                    </Modal.Body>
                    <Modal.Footer>
                        <Button color="gray" onClick={() => setOpenModalDetil(false)}>
                            Tutup
                        </Button>
                    </Modal.Footer>
                </Modal>
            </form>

            <form key={'form-edit'}>
                <Modal id='modal-edit' show={openModalEdit} onClose={() => setOpenModalEdit(false)}>
                    <Modal.Header>Edit Task #{editTask.id_task}</Modal.Header>
                    <Modal.Body>
                        <div className="space-y-3">
                            <div>
                                <div className="mb-2 block">
                                    <Label htmlFor="judul" value="Judul" />
                                </div>
                                <TextInput
                                    {...regFormEdit("title", {
                                        required: "Judul wajib diisi",
                                    })}
                                    id="title" placeholder="Masukan judul" required />
                                <ul>
                                    {errorFormEdit.title ? <li className='TextError'>{errorFormEdit.title.message} </li> : null}
                                </ul>
                            </div>
                            <div>
                                <div className="mb-2 block">
                                    <Label htmlFor="deskripsi" value="Deskripsi" />
                                </div>
                                <Textarea
                                    {...regFormEdit("description", {
                                        required: "Deskripsi wajib diisi",
                                    })}
                                    id="description" rows={4} placeholder="Masukan deskripsi" required />
                                <ul>
                                    {errorFormEdit.description ? <li className='TextError'>{errorFormEdit.description.message} </li> : null}
                                </ul>
                            </div>
                        </div>

                    </Modal.Body>
                    <Modal.Footer>
                        <Button onClick={handleSubmitUpdate(submitUpdateTask)}>Simpan</Button>
                        <Button color="gray" onClick={() => setOpenModalEdit(false)}>
                            Batal
                        </Button>
                    </Modal.Footer>
                </Modal>
            </form>

            <form key={'form-delete'}>
                <Modal id='modal-delete'  show={openModalDelete} size="md" onClose={() => setOpenModalDelete(false)} popup>
                    <Modal.Header />
                    <Modal.Body>
                        <div className="text-center">
                            <HiOutlineExclamationCircle className="mx-auto mb-4 h-14 w-14 text-gray-400 dark:text-gray-200" />
                            <h3 className="mb-5 text-lg font-normal text-gray-500 dark:text-gray-400">
                                Anda yakin ingin menghapus task ini?
                            </h3>
                            <div className="flex justify-center gap-4">
                                <Button color="failure" onClick={() => submitDeleteTask()}>
                                    Ya, Hapus
                                </Button>
                                <Button color="gray" onClick={() => setOpenModalDelete(false)}>
                                   Batal
                                </Button>
                            </div>
                        </div>
                    </Modal.Body>
                </Modal>
            </form>
        </div>
    )
}
