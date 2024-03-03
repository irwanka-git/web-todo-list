import React from 'react'
import { Outlet } from 'react-router-dom';
import { Header } from './Header';
import {Helmet, HelmetProvider} from "react-helmet-async"

import { useRecoilValue } from "recoil";
import { pagetTitleState } from '../state';

export const BaseLayout = () => {
    const pageTitle = useRecoilValue(pagetTitleState)

    return (
        <HelmetProvider>
            <Helmet>
                <title>{pageTitle ? pageTitle : "No title"}</title>
            </Helmet>
            <div className='flex-row max-h-screen overflow-auto mx-2 my-2'>
                <Header />
                <main className={`mx-2 my-3 mt-6 text-sm`}>
                    <Outlet />
                </main>
            </div>
        </HelmetProvider>
    )
}
