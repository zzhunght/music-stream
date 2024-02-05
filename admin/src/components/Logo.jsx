
import Image from 'next/image';
import logo from '../assets/images/logo.jpg'

const Logo = () => {
    return <div>
        <Image src={logo} className='w-[54px] h-[54px]' alt='logo'/>
    </div>
}

export default Logo;