Quite possibly the worlds most easiest blogging platform.

Design goals and rough tasks:

-  Serves blog site
-  Background daemon pulls posts from github repo.
-  Encodes html from markdown.  Word count, reading time, gravatar, etc...
-  Person only has comming a markdown file to their repo.  boom.  done.
-  Site should be responsive out of the box (skeleton?)
-  Need to make sure that comments embedded in markdown don't have issues with rendering engine.
-  Provides basic scaffolding for blog site (tag cloud, navigation, etc..).  But completely overridable.
-  Overrides should be done in yaml file.  e.g. location of repo.  name.  etc...
-  Include blurb at bottom of site, "Site delivered by scribl"  with a link to the OSS repo.
-  Add apache license to repo.
-  Add issues and project page to scribl repo.
-  Be unoppionated about how you host the site.  
-  For my personal site (yawn.me) front with traefik ?
-  Should I bake ssl into scribl with let's encrypt ?
