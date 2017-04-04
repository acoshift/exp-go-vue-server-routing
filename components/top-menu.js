Vue.component('top-menu', {
  template: `
    <div>
      This is Top-Menu<br>
      <a href="/">Index</a><br>
      <a href="/list">List</a><br>
    </div>
  `,
  methods: {
    signIn () {
      window.alert('sign in clicked')
    }
  }
})
