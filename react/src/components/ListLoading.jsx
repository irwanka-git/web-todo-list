 import React from 'react'
import Skeleton, { SkeletonTheme } from 'react-loading-skeleton'
import 'react-loading-skeleton/dist/skeleton.css'

export const ListLoading = () => {
    return (
         
            <SkeletonTheme baseColor='rgb(226 232 240)' highlightColor="rgb(203 213 225)">
                <Skeleton height={30} className='mt-2'/>
                <Skeleton height={30} className='mt-2'/>
                <Skeleton height={30} className='mt-2'/>
            </SkeletonTheme>
         
    )
}
