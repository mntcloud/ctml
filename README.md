<h1 align="center"> ConvenienT Markup Language</h1>

It was created as a solution to my problems in writing HTML pages, let me show the format itself:
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
What you saw above that it is all syntax (yet).
In result it transforms in something like this:
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

## What things does it have right now?

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

