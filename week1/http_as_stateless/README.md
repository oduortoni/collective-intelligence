### HTTP and Statelessness

#### Stateless?

HTTP is stateless, meaning each request you make is independent. For example, if you visit `/about` and then `/blog`, the server has no memory of your first request when processing the second one. This is how HTTP is designed—it doesn't keep track of previous requests.

#### Why Make It Stateful?

Imagine trying to access a `/dashboard` that requires you to be logged in. Without remembering your login status, you'd have to log in again every time you try to access a protected route. This would be repetitive and frustrating. To improve user experience, HTTP allows ways to "remember" the client’s state. One such way is through **cookies**. Cookies store small bits of information that help the server recognize the user, so they don't have to log in repeatedly.

#### What We'll Cover

In this example, we’ll implement a simple way for the server to remember the client using cookies. The server will generate a token and store it in a CSV file. After logging in, the server sends the token to the browser as a cookie. The browser will then send that cookie with every subsequent request, allowing the server to remember the client.

This will help you understand how HTTP's stateless nature works and how HTTP headers can be used to manage client-server communication.

###