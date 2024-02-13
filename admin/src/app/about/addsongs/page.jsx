import TitlePage from "@/components/titlepage/TitlePage";
import { ArrowUpToLine } from "lucide-react";
import CsInputFile from "./cs-input-file";
import { Label } from "@/components/ui/label";
import { Input } from "@/components/ui/input";
import SaveCancel from "@/components/savecancel/SaveCancel";

const page = () => {
    return (
        <div >
            <TitlePage title={'Add Songs'}/>
            <div className="flex flex-col">
                <div className="flex justify-between h-[200px] space-x-20">
                    <CsInputFile text={'Audio file'} note={'MP3, or WOFF2,(MAX 5MB)'}/>
                    <CsInputFile text={'Thumbnail'} note={'JPG, PNG,(MAX 5MB)'}/>
                </div>

                <div className="mt-10 space-y-6">
                    <div className="grid w-full items-center gap-1.5">
                        <Label htmlFor="lname">Song Name</Label>
                        <Input type="lname" id="lname" placeholder="Last name" />
                    </div>

                    <div className="grid w-full items-center gap-1.5">
                        <Label htmlFor="lname">Artists</Label>
                        <Input type="lname" id="lname" placeholder="Last name" />
                    </div>

                    <div className="grid w-full items-center gap-1.5">
                        <Label htmlFor="lname">Lyrics</Label>
                        <Input type="lname" id="lname" placeholder="Last name" className='h-60'/>
                    </div>
                </div>

                <div className="mt-10">
                    <SaveCancel/>
                </div>
            </div>
        </div>
    )
}

export default page;