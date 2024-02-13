import SaveCancel from "../savecancel/SaveCancel";
import { Button } from "../ui/button";
import { Input } from "../ui/input";
import { Label } from "../ui/label";

const FormProfile = () => {

    return (
        <form className="space-y-5">
            <div className="flex flex-col space-y-5">
                <div className="flex justify-between items-center space-x-8">
                    <div className="grid w-full max-w-xl items-center gap-1.5">
                        <Label htmlFor="fname">First Name</Label>
                        <Input type="fname" id="fname" placeholder="First name" />
                    </div>
                    <div className="grid w-full max-w-xl items-center gap-1.5">
                        <Label htmlFor="lname">Last Name</Label>
                        <Input type="lname" id="lname" placeholder="Last name" />
                    </div>
                </div>

                <div className="grid w-full items-center gap-1.5">
                    <Label htmlFor="email">Email</Label>
                    <Input type="email" id="email" placeholder="Email" />
                </div>

                <div className="grid w-full items-center gap-1.5">
                    <Label htmlFor="address">Address</Label>
                    <Input type="address" id="address" placeholder="Address" />
                </div>

                <div className="grid w-full items-center gap-1.5">
                    <Label htmlFor="cnumber">Contact number</Label>
                    <Input type="cnumber" id="cnumber" placeholder="Contact number" />
                </div>

                <div className="flex justify-between items-center space-x-8">
                    <div className="grid w-full max-w-xl items-center gap-1.5">
                        <Label htmlFor="city">City</Label>
                        <Input type="city" id="city" placeholder="City" />
                    </div>
                    <div className="grid w-full max-w-xl items-center gap-1.5">
                        <Label htmlFor="state">State</Label>
                        <Input type="state" id="state" placeholder="State" />
                    </div>
                </div>

                <div className="grid w-full items-center gap-1.5">
                    <Label htmlFor="password">Password</Label>
                    <Input type="password" id="password" placeholder="Password" />
                </div>
            </div>
            <SaveCancel/>
        </form>
    )
}

export default FormProfile;