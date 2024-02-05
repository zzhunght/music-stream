'use client';

import { ChevronDown } from "lucide-react";
import { usePathname, useRouter } from "next/navigation";
import { useMemo, useState } from "react";
import SubmenuItem from "./SubmenuItem";

const SidebarItem = props => {
    const { name, path, icon: Icon, items } = props.item;
    const [expanded, setExpanded] = useState(false)

    const router = useRouter();
    const pathname = usePathname();

    const onClick = () => {
        if(items && items.length > 0) {
            return setExpanded(!expanded)
        }
        router.push(path);
    }

    const isActive = useMemo(() => {
        if(items && items.length > 0) {
            if(items.find(item => item.path === pathname)) {
                setExpanded(true)
                return true
            } 
            // else {
            //     setExpanded(false)
            //     return false
            // }
        }
        return path === pathname
    }, [path, pathname])

    return (
        <>
            <div 
                className={`flex items-center p-3 pl-5 hover:bg-sidebar-background cursor-pointer hover:text-sidebar-active justify-between ${isActive && "text-sidebar-active bg-sidebar-background"}`}
                onClick={onClick}
            >
                <div className="flex items-center space-x-2">
                    <Icon size={28}/>
                    <p className={`text-sm ${isActive && "font-semibold"}`}>
                        {name}
                    </p>
                </div>
                {items && items.length > 0 && (
                    <ChevronDown size={18}
                        className={expanded ? "rotate-180 duration-200" : ""}
                    />
                )}
            </div>
    
            {expanded && 
                items &&
                items.length > 0 &&
                <div className="flex flex-col space-y-4 ml-10">
                    {items.map(item => (
                        <SubmenuItem key={item.path} item={item}/>
                    ))}
                </div>
            }
        </>
    )
}

export default SidebarItem;