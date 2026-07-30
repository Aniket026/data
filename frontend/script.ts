const uploadBtn = document.getElementById("uploadBtn") as HTMLButtonElement;

const oldFileInput = document.getElementById("oldFile") as HTMLInputElement;

const newFileInput = document.getElementById("newFile") as HTMLInputElement;

const result = document.getElementById("result") as HTMLDivElement;

uploadBtn.addEventListener("click", async () => {

    if (
        !oldFileInput.files?.length ||
        !newFileInput.files?.length
    ) {
        alert("Please select both files.");
        return;
    }

    const formData = new FormData();

    formData.append(
        "oldFile",
        oldFileInput.files[0]
    );

    formData.append(
        "newFile",
        newFileInput.files[0]
    );

    try {

        result.innerHTML = "Uploading...";

        const response = await fetch(
            "http://localhost:8080/upload",
            {
                method: "POST",
                body: formData
            }
        );

        const data = await response.json();

        if(response.ok){

            result.innerHTML=`

<b>Success</b>

<br><br>

Old File :

${data.oldFile}

<br>

New File :

${data.newFile}

`;

        }else{

            result.innerHTML=data.error;

        }

    } catch (err) {

        result.innerHTML="Server Error";

        console.error(err);

    }

});