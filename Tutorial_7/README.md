# Tutorial 7 : Model loading

Until now, we hardcoded our cube directly in the source code. I’m sure you will agree that this was cumbersome and not very handy.

In this tutorial we will learn how to load 3D meshes from files. We will do this just like we did for the textures : we will write a tiny, very limited loader, and I’ll give you some pointers to actual libraries that can do this better that us.

To keep this tutorial as simple as possible, we’ll use the OBJ file format, which is both very simple and very common. And once again, to keep things simple, we will only deal with OBJ files with 1 UV coordinate and 1 normal per vertex (you don’t have to know what a normal is right now).
