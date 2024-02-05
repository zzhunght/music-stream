'use client'
import React, { useState, useEffect} from "react";
import ThemeToggler from "./ThemeToggler";
import Nav from "./Nav";
import MobileNav from "./MobileNav";

const Header = () => {

    return (
        <header className="h-[80px] flex items-center">
            <div className="container mx-auto">
                <div className="flex justify-end items-center">
                    <div className="flex items-center gap-x-6">
                        {/* Nav */}
                        <Nav containerStyles='hidden xl:flex gap-x-8 items-center'/>
                        <ThemeToggler/>
                        <div className="xl:hidden">
                            <MobileNav/>
                        </div>
                    </div>
                </div>
            </div>
        </header>
    )
}

export default Header;