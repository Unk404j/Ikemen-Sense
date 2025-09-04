## Summary

- Added explicit `windows && sdl` build tags and SDL2 linkage directives  
  → Enables the Windows-only SDL input backend with HIDAPI hints and haptic fallback  

- Introduced a **stubbed implementation** for non-SDL builds  
  → Routed Lua script initialization through a new `registerRumble` helper with build-aware registration  

- Updated Lua helpers and documentation  
  → Accepts **intensity** and **duration** parameters for vibrations, reflected in both the preset module and README build instructions  

---

## Testing

```bash
⚠️ go build -o /tmp/IkemenSense ./src
# /usr/bin/ld: cannot find -lXxf86vm: No such file or directory
```

## Notes

```bash
Building still requires additional system libraries
→ Install libXxf86vm-dev (or equivalent) to satisfy linker dependencies
