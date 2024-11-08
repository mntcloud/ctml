# Convenient Markup Language
A format for building and maintaining varied websites 

## Format Example
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
- [ ] Better developer UX with reload-on-changes server out of the box
- [ ] Comments (yea, they're not implemented)
- [ ] Subsyntax for marking text parts italic, bold or even including inside of a span with keeping text readability
- [ ] Repetative elements with an option of exclusion
- [ ] Integrating templating as a part of language design itself

# Install 
TODO:
