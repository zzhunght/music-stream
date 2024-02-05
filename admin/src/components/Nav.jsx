import { BellIcon } from "lucide-react";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "./ui/select";
import Image from "next/image";
import logo from '../assets/images/logo.jpg'

const Nav = ({containerStyles, linkStyles, underlineStyles}) => {
    return <nav className={`${containerStyles}`}>
        <div>
            <Select>
                <SelectTrigger className='w-[180px]'>
                    <Image src={logo} className="w-[30px]"/>
                    <SelectValue placeholder={'Nhóm 12'}/>
                </SelectTrigger>
                <SelectContent>
                    <SelectItem value="light">Light</SelectItem>
                    <SelectItem value="dark">Dark</SelectItem>
                    <SelectItem value="system">System</SelectItem>
                </SelectContent>
            </Select>
        </div>
        <BellIcon/>
        
    </nav>
}

export default Nav;