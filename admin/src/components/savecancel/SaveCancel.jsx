const { Button } = require("../ui/button")

const SaveCancel = () => {

    return(
        <div className="flex items-start space-x-6">
            <Button variant="destructive">
                Cancel
            </Button>
            <Button variant="save">
                Save
            </Button>
        </div>
    )
}

export default SaveCancel;