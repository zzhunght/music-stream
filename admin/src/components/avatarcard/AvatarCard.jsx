'use client'

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

const AvatarCard = () => {

    return (
        <Dialog>
            <div className="h-full w-full flex justify-center items-center ">
                <DialogTrigger>
                    <div className="cursor-pointer hover:bg-sidebar-background flex justify-between items-center w-full h-full">
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
            </div>
            <DialogContent>
            <DialogHeader>
                <DialogTitle>Are you absolutely sure?</DialogTitle>
                <DialogDescription>
                This action cannot be undone. This will permanently delete your account
                and remove your data from our servers.
                </DialogDescription>
            </DialogHeader>
            </DialogContent>
        </Dialog>
    )
}

export default AvatarCard;