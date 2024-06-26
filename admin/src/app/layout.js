import { Inter } from "next/font/google";
import "./globals.css";
import Header from "@/components/Header";
import Footer from "@/components/Footer";
import ThemeProvider from "@/components/ThemeProvider";
import Sidebar from "@/components/sidebar/Sidebar";

const inter = Inter({ subsets: ["latin"] });

export const metadata = {
  title: "Create Next App",
  description: "Generated by create next app",
};

export default function RootLayout({ children }) {
  return (
    <html lang="en">
      <body className={inter.className}>
          <ThemeProvider attribute='class' defaultTheme='light'>
            <div className="flex h-screen w-full">
              <Sidebar/>
              <div className="flex flex-col w-full h-full ml-64 p-4">
                {/* <Header/> */}
                <div className="p-10">{children}</div>
                <Footer/>
              </div>
            </div>
          </ThemeProvider>
        </body>
    </html>
  );
}
