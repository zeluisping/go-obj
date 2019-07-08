# go-obj
A pure Go implementation of a Wavefront .obj loader

# todo
## Vertex data
- rational and non-rational forms of curve or surface type:
    basic matrix, Bezier, B-spline, Cardinal, Taylor (cstype)
-  degree (deg)
- basis matrix (bmat)
- step size (step)

## Elements
- point (p)
- line (l)
- curve (curv)
- 2D curve (curv2)
- surface (surf)

## Free-form curve/surface body statements
- parameter values (param)
- outer trimming loop (trim)
- inner trimming loop (hole)
- special curve (scrv)
- special point (sp)
- end statement (end)

## Connectivity between free-form surfaces
- connect (con)

## Grouping
- group name (g)
- smoothing group (s)
- merging group (mg)
- object name (o)

## Display/render attributes
- bevel interpolation (bevel)
- color interpolation (c_interp)
- dissolve interpolation (d_interp)
- level of detail (lod)
- material name (usemtl)
- material library (mtllib)
- shadow casting (shadow_obj)
- ray tracing (trace_obj)
- curve approximation technique (ctech)
- surface approximation technique (stech)

## General statement
- call filename.ext arg1 arg2
