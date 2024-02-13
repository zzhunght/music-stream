import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { ArrowUpToLine } from "lucide-react";

const CsInputFile = props => {

    return(
           <div className="w-full flex flex-col">
                <Input id='picture' type='file' className='invisible'/>
                <div className="w-full h-full border-dashed border-2 border-gray-300 rounded-lg">
                    <Label htmlFor="picture" className='cursor-pointer h-full w-full'>
                        <div className="flex flex-col justify-center items-center space-y-5 h-full">
                            <ArrowUpToLine/>
                            <div>
                                <span className="flex justify-center font-semibold">
                                    {props.text ? props.text : 'File'}
                                </span>
                                <p className="flex justify-center text-sm text-blur">
                                    {props.note ? props.note : 'Note'}
                                </p>
                            </div>
                        </div>
                    </Label>
                </div>
           </div>
    )
}

export default CsInputFile;