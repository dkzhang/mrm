// <script>
    document.querySelectorAll('.cell-occupied').forEach(rect => {
    rect.addEventListener('click', (event) => {
        // Stop the event from bubbling up to the SVG
        event.stopPropagation();

        // Get the new rectangle's text
        let TextName = document.getElementById('textName');
        let TextApplicant = document.getElementById('textApplicant');

        // Set the text of the new rectangle's text
        TextName.textContent = rect.dataset.name; // Get the name from the "data-name" attribute
        TextApplicant.textContent = rect.dataset.applicant;

        // Compute the width of the new rectangle based on the length of the text
        let newRectWidth = TextName.getComputedTextLength() + 20; // Add some padding

        // Get the center of the clicked rectangle
        let rectCenter = parseFloat(rect.getAttribute('x')) + parseFloat(rect.getAttribute('width')) / 2;

        // Calculate the new x position for the new rectangle
        let newRectX = rectCenter - newRectWidth / 2;
        let newRectY = rect.getAttribute('y') - 60; // Position it above the clicked rectangle

        // Get the new rectangle
        let newRect = document.getElementById('newRect');

        // Set the width of the new rectangle
        newRect.setAttribute('width', newRectWidth);

        // Translate the new rectangle and its text
        newRect.setAttribute('transform', "translate(${newRectX} ${newRectY})");
        TextName.setAttribute('transform', "translate(${newRectX + 10} ${newRectY + 20})"); // Add some padding
        TextApplicant.setAttribute('transform', "translate(${rectCenter} ${newRectY + 40})"); // Add some padding

        // Make the new rectangle and its text visible
        newRect.style.visibility = 'visible';
        TextName.style.visibility = 'visible';
        TextApplicant.style.visibility = 'visible';
    });
});

    // Add an event listener to the SVG
    document.getElementById('mySvg').addEventListener('click', () => {
    // Hide the new rectangle and its text
    document.getElementById('newRect').style.visibility = 'hidden';
    document.getElementById('newRectText').style.visibility = 'hidden';
    document.getElementById('textName').style.visibility = 'hidden';
});
// </script>