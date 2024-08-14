document.addEventListener('DOMContentLoaded', function () {
    const form = document.getElementById('note-form');
    const notesList = document.getElementById('notes-list');

    function fetchNotes() {
        fetch('/notes')
            .then(response => response.json())
            .then(data => {
                notesList.innerHTML = '';
                data.forEach(note => {
                    const noteElement = document.createElement('div');
                    noteElement.classList.add('note');
                    noteElement.innerHTML = `
                        <div class="note-title">${note.title}</div>
                        <div class="note-content">${note.content}</div>
                        <div class="note-date">${new Date(note.created_at).toLocaleString()}</div>
                    `;
                    notesList.appendChild(noteElement);
                });
            })
            .catch(error => console.error('Error fetching notes:', error));
    }

    form.addEventListener('submit', function (e) {
        e.preventDefault();

        const title = document.getElementById('title').value;
        const content = document.getElementById('content').value;

        fetch('/notes', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ title, content })
        })
            .then(response => response.json())
            .then(data => {
                form.reset();
                fetchNotes();
            })
            .catch(error => console.error('Error adding note:', error));
    });

    fetchNotes();
});
