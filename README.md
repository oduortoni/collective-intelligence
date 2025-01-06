# Topics in Programming

## Description

This repo aims to provide a simple, hands-on introduction to key programming concepts. We’ll focus on fundamental ideas that are essential when building applications. While frameworks might play a role in your future work, understanding the basics is crucial since frameworks are built on these core principles. Frameworks may come and go, but the foundational concepts remain constant. By mastering the basics, you’ll find it easier to learn more advanced topics, which are often just combinations of these basics.

AI will play a role in the future, but with the right knowledge, you can use it as a helpful tool, not just a solution provider. How will you know if AI’s code is optimized if you don’t understand the fundamentals? And when AI struggles, will you be able to break the problem down into smaller parts to help it succeed?

## Topics

1) HTTP as a Stateless protocol

What does "stateless" mean? If each request is independent, how does the server know a user is logged in instead of prompting them to log in every time? What are sessions or cookies? We will try to implement a very simple authentication mechanism that uses a file to remember a user's state each time they make a request. We will prove that even though each request may be independent and the server does not remember the user's previous actions, we can implement some features on top of http to make it easy to remember a user's actions.
