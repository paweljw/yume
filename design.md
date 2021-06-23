# Design document

## Table of contents

%TOC%

## Text or graphics?

*Text*. We may want to go with a custom client at some point, support [MCP](https://www.moo.mud.org/mcp/mcp2.html) commands, even provide our own packages for MCP-compatible clients, but the initial goal is "if it runs telnet, it plays". Furthermore, initial playtesting is encouraged to be done through telnet and not a dedicated MUD client.

## Rooms or coordinates?

*Rooms*. Designing for in-room coordinates is much too big of a hassle, and at this stage is probably unnecessary. Indeed it may never be necessary.

## Scripting

Current idea is to script as much as possible. Certain actions will have to be implemented in the core, but when deciding "is this a core addition or a script", lean "script".

### Examples of actions that should be in the core

* Sign-in process
* Character creation process
* Movement (moving between rooms)
* Attacking with melee weaponry
* Casting (as a command, `cast` - though the spell will be scripted)
* Using objects
* Wearing items
* Picking up/dropping items
* If ever added, attacking with range weaponry

### Examples of things that should be scripts

* A spell
  * Typical spell will take a `caster` and a `target` object (these may be players, monsters, in-game objects or any combination thereof). The spell should check whether it can be cast by the `caster`, does it apply to the `target` and so on
* A locked door
  * A locked door's script can accept an `interactee`, which will usually be a player, and check whether they have a specific flag. If so, the script allows passage and the core executes the movement
* An object, like a key
  * A given object such as a key, may have its own script. Said script will check, upon being invoked by `use`, it's `target` and if it has the correct ID it will set a flag on the player (that the door is now open).
* A monster's AI
  * Monsters' behaviors should be scripted also. This will allow for quite some flexibility in case of more complex, multi-stage, or "boss" enemies.

## Resets / repops

Monsters will have a spawn location (or a list of them) associated with them. They will also have a timer (which may be a formula?), expressed in ticks (server tick will probably equal 1s for simplicity). Once the monster is killed, its presence is removed from the map, but the monster "object", in its dead state, is retained. The game checks the list of monsters anyway - to make them execute the next step in their script - so once enough ticks have passed, the monster is repopulated into one of their spawn locations.

Gathering nodes (mining, herbalism, fishing and so on) will work in a similar fashion - the node itself is persistent, but is not apparent to the players once exhausted. These nodes shall also be checked periodically whether they need to repop.

## Combat

### Melee combat

* Auto-attack: while in combat, auto-attacks happen every so often (this will fall under both weapon speed and dexterity probably). These are so-called "white damage" - it ain't much, but it's honest work. These attacks just keep happening until the combat is done.
* Weaponskills: while in combat, a weaponskill (a specific command, such as "slash" or "jab") may be executed. These will put the player into a "casting" state for a set amount of ticks (or maybe seconds can be viable?) (probably modified by dexterity again). If the character takes any other action while in the "casting" period, it will be ignored. Weaponskills may have cooldowns of X ticks, such as spells. Weaponskills consume stamina.

Both auto attacks and weaponskills can be:
* dodged - based on the opponent's dex
* parried - based on the opponent's dex, strength, luck?
* shrugged off (think "glancing blow") - if the disparity in the "body" level is too high

### Magic combat



