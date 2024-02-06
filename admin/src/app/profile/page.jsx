import FormProfile from "@/components/formprofile/FormProfile";
import TitlePage from "@/components/titlepage/TitlePage";
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";

const Profile = () => {
    return (
        <div className="h-screen">
            <TitlePage title={'Edit profile'}/>
            <div className="flex justify-end items-center space-x-10">
                <Input id='picture' type='file' className='invisible'/>
                <Label htmlFor="picture" className='h-10 px-4 py-2 border border-gray-500 rounded-md flex items-center cursor-pointer'>
                    Change
                </Label>

                <Avatar className="h-[80px] w-[80px]">
                    <AvatarImage src={"https://github.com/shadcn.png"} />
                    <AvatarFallback>CN</AvatarFallback>
                </Avatar>

            </div>
            <FormProfile/>
        </div>
    )
}

export default Profile;