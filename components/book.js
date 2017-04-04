Vue.component('book', {
  template: `
    <div>
      <h3>Book Detail</h3>
      ID: {{ book.id }}<br>
      Title: {{ book.title }}<br>
      Description: {{ book.description }}
    </div>
  `,
  props: {
    book: Object
  }
})
