# Convenient Markup Language
A format for building and maintaing varied websites 
 
## Goal
Project was created to resolve my problems with building HTML pages, like:
- repetitiveness
    - I know about HTML-template engines. I don't like it too much because of implementation and use
      - Engines exists mostly like libraries, than standalone programs.
          - And templating format is not so pretty at micro and major scale, looks like something foreign, better be deeply intergrated 
      - Standalone programs, that include engines, I don't like their focus on something one (blog or whatever) and they put some restrictions.
          - Their's documentation most likely is horrible, this what i've encountered so far
      - Conclusion: I would like to have a middle of these things, a flexible solution for different use cases.
- development
    - code edit + manual browser reload != happiness

Also I wanted to make my own handwritten parser and lexer, that is reason why it has new syntax. 
Other, I wanted to improve perfomance of writing webpages.

# Format Example
```
<body
    <header .class1 .class2
        <button: Home
        <button =attr false: Pages
        <button =attr false: About
    <main .class3 .class4
        <h1: CTML
        <p
            Format for writing HTML pages in a harmony
            Pretty cool, right?
    <footer
        made by mntcloud
```
> This page contains all the syntax, that is defined in the program.
  When you know HTML, it looks almost intuitive, don't you think?

Then the program translates the CTML page above into this HTML:
```html
<html>
    <!-- ... -->
    <body>
        <header class="class1 class2">
            <button>Home</button>
            <button attr="false">Pages</button>
            <button attr="false">About</button>
        </header>
        <main class="class3 class4">
            <h1>CTML</h1>
            <p> 
                Format for writing HTML pages in a harmony
                Pretty cool, right? 
            </p>
        </main>
        <footer>
            made by mntcloud
        </footer>
    </body>
</html>
```

## What features does it have right now?

- [x] Intuitive syntax, that has similiraties with HTML, but with Python identation 
- [x] Swapping tags names easily any without help of third-party extensions, thanks to the syntax
- [x] Assigning classes to element without declaring class attribute
- [x] Better developer UX with reload-on-changes server out of the box
- [ ] Comments (yea, they're not implemented)
- [ ] Subsyntax for marking some text parts italic, bold or even including inside of a span with keeping text readability almost the same
- [ ] Repetative elements with an option of exclusion
- [ ] Integrating templating as a part of language design itself
- [ ] Mode for working with CSS and JS in one file, like Svelte does, but with some design differences

# Install 

