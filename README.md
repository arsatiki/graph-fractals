# graph-fractals

Compute fractal dimension of OSM bike maps.

# TODO

- [ ] Parse OSM XML files (see http://wiki.openstreetmap.org/wiki/OSM_XML)
- [ ] Extract bike path graph
- [ ] Compute CBB for various box sizes (see http://arxiv.org/pdf/cond-mat/0701216v1.pdf)
  - [ ] Normalize distances to meters
  - [ ] Implement A* (with cutoff)
  - [ ] Calculate average box sizes instead of cutoff
- [ ] Estimate the fractal dimension (see http://mathworld.wolfram.com/LeastSquaresFittingPowerLaw.html)
