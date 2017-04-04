Vue.component('books', {
  template: `
    <ul>
      <li v-for="x in list">
        <a :href="'/view?id='+x.id">{{ x.title }}</a>
      </li>
    </ul>
  `,
  props: {
    list: Array
  }
})
