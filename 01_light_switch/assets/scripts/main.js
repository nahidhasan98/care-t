console.log("script connected successfully.");

let bulbHead = document.getElementById('bulbHead');
let bulbNeck = document.getElementById('bulbNeck');
let bulbTail = document.getElementById('bulbTail');

// taking all bulb elements to an array
let bulbElements = [bulbHead, bulbNeck, bulbTail]

// initially light off
let bulbStatus = false;

// light on/off function
bulbElements.forEach((item) => {
    item.addEventListener("click", () => {
        if (bulbStatus) {   // already on. turning off
            bulbHead.style.background = "#fff";
            bulbNeck.style.background = "#f1f1f1";

            bulbStatus = false;
        } else {    // already off. turning on
            bulbHead.style.background = "#ffe666";
            bulbNeck.style.background = "#ffe666";

            bulbStatus = true;
        }
    })
});


// light on/off (image section)
let bulbImage = document.getElementsByClassName('bulbImage');
let offImage = document.getElementById('offImage');
let onImage = document.getElementById('onImage');
let imageBulbStatus = false;

bulbImage[0].addEventListener("click", function () {
    if (imageBulbStatus) {   // already on. turning off
        offImage.style.display = "";
        onImage.style.display = "none";

        imageBulbStatus = false;
    } else {    // already off. turning on
        offImage.style.display = "none";
        onImage.style.display = "";

        imageBulbStatus = true;
    }
});