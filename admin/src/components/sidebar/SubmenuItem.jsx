'use client';

import { usePathname, useRouter } from "next/navigation";
import React, { useMemo, useState } from "react";

const SubmenuItem = props => {
    const { name, path } = props.item;

    const router = useRouter();
    const pathname = usePathname();

    const onClick = () => {
        router.push(path);
    }

    const isActive = useMemo(() => {
        return path === pathname
    }, [path, pathname])
    return <div className={`text-sm hover:text-sidebar-active hover:font-semibold cursor-pointer ${isActive && "text-sidebar-active font-semibold"}`}
        onClick={onClick}
    >
        {name}
    </div>
}

export default SubmenuItem;