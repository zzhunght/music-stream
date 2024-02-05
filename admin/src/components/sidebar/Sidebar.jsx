'use client';

import { LayoutDashboard, Disc, User, UserCog, LayoutTemplate, ScatterChart, Settings } from 'lucide-react'
import SidebarItem from './SidebarItem'
import AvatarCard from '../avatarcard/AvatarCard';

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
    
]

const advancedMenus = [
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

const personalMenus = [
    {
        name: 'Settings',
        path: '/settings',
        icon: Settings,
    },
]

const Sidebar = () => {



    return (
        <div className='fixed top-0 left-0 bg-[#ECEEF2] w-64 h-screen shadow-lg z-10'>
            <div className='flex flex-col space-y-5'>
                <div className='w-full flex justify-between items-center h-[80px]'>
                    <AvatarCard/>
                </div>
                
                <div className='flex flex-col space-y-1'>
                    <p className='text-xs ml-5 text-sidebar-color'>
                        Common
                    </p>
                    {itemsMenu.map((item) => (
                        <SidebarItem key={item.path} item={item}/>
                    ))}
                    <p className='text-xs ml-5 text-sidebar-color'>
                        Advanced
                    </p>
                    {advancedMenus.map((item) => (
                        <SidebarItem key={item.path} item={item}/>
                    ))}
                    <p className='text-xs ml-5 text-sidebar-color'>
                        Personal
                    </p>
                    {personalMenus.map((item) => (
                        <SidebarItem key={item.path} item={item}/>
                    ))}
                </div>
            </div>
        </div>
    )
}

export default Sidebar;