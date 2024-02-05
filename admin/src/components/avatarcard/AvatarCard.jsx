'use client'

import { useRouter } from 'next/navigation';
import { Avatar, AvatarImage } from "../ui/avatar";
import { AvatarFallback } from '@radix-ui/react-avatar';
import { ChevronRight } from 'lucide-react'
import {
    Dialog,
    DialogContent,
    DialogDescription,
    DialogHeader,
    DialogTitle,
    DialogTrigger,
  } from "@/components/ui/dialog"

const items = [
    {
        name: 'Edit Profile',
        path: '/admin/profile',
        icon: ChevronRight
    },
]

const AvatarCard = () => {
    const router = useRouter();

    const handleClick = () => {
        router.push('/profile')
    }
    return (
        <div className="h-full w-full flex justify-center items-center cursor-pointer hover:bg-sidebar-background">
            <Dialog>
                    <DialogTrigger>
                        <div className="flex justify-between items-center w-full h-full">
                            <div className='flex items-center space-x-2'>
                                <Avatar>
                                    <AvatarImage src={"https://github.com/shadcn.png"} />
                                    <AvatarFallback>CN</AvatarFallback>
                                </Avatar>
                                <p>
                                    Tran Minh Hoang
                                </p>
                            </div>
                            <ChevronRight size={18}/>
                        </div>
                    </DialogTrigger>
                <DialogContent>
                    <DialogHeader>
                        {/* <DialogTitle>Are you absolutely sure?</DialogTitle> */}
                        <DialogDescription>
                            <>
                                {items.map(item => (
                                    <div 
                                        key={item.path} 
                                        className="flex justify-between items-center cursor-pointer hover:bg-sidebar-background h-[40px] rounded-md"
                                        onClick={handleClick}
                                        >
                                        <p className="text-black ml-2">
                                            {item.name}
                                        </p>
                                        <ChevronRight/>
                                    </div>
                                ))}
                                <div className="flex justify-between items-center cursor-divointer hover:bg-sidebar-background h-[40px] rounded-md">
                                    <p className="text-[#FF6B6B] ml-2">Logout</p>
                                </div>
                            </>
                        </DialogDescription>
                    </DialogHeader>
                </DialogContent>
            </Dialog>
        </div>
    )
}

export default AvatarCard;