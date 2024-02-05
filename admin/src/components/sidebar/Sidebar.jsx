'use client';

import Logo from '../Logo'

import { LayoutDashboard, Disc, User, UserCog, LayoutTemplate, ScatterChart } from 'lucide-react'
import SidebarItem from './SidebarItem'

const itemsMenu = [
    {
        name: 'Dashboard',
        path: '/',
        icon: LayoutDashboard
    },
    {
        name: 'Songs & Related',
        path: '/about',
        icon: Disc,
        items: [
            {
                name: 'Songs',
                path: '/about'
            },
            {
                name: 'Authors',
                path: '/about/authors'
            },
            {
                name: 'Categories',
                path: '/about/categories'
            },
            {
                name: 'Albums',
                path: '/about/albums'
            },
            {
                name: 'Playlists',
                path: '/about/playlists'
            },
        ]
    },
    {
        name: 'Account',
        path: '/account',
        icon: User
    },
    {
        name: 'Employee',
        path: '/employee',
        icon: UserCog
    },
    {
        name: 'Role',
        path: '/role',
        icon: LayoutTemplate
    },
    {
        name: 'Analytics',
        path: '/analytics',
        icon: ScatterChart
    },
]

const Sidebar = () => {
    return (
        <div className='fixed top-0 left-0 bg-[#ECEEF2] w-64 h-screen shadow-lg z-10'>
            <div className='flex flex-col space-y-10'>
                <div className='w-full flex justify-center items-center'>
                    <Logo/>
                </div>
                
                <div className='flex flex-col space-y-1'>
                    {itemsMenu.map((item) => (
                        <SidebarItem key={item.path} item={item}/>
                    ))}
                </div>
            </div>
        </div>
    )
}

export default Sidebar;