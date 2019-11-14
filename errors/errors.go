package errors

import (
  goErrors "errors"
)

var DIR_IS_NOT_A_DIRECTORY = goErrors.New( "directory is not a directory ;-)")
var DIR_DOES_NOT_EXIST = goErrors.New( "directory does not exist")
var DATADIR_IS_LOCKED = goErrors.New( "data directory is locked" )
var DATADIR_IS_INVALID = goErrors.New( "invalid data directory" )
var NO_SUCH_SOURCE_TYPE = goErrors.New( "no such source type" )
var NO_SUCH_SOURCE = goErrors.New( "no such source" )
var DUPLICATE_SOURCE = goErrors.New( "duplicate source" )
var SOURCE_ADD_NO_SOURCE = goErrors.New( "source add: no source" )
var SOURCE_DELETE_NO_SOURCE = goErrors.New( "source delete: no source" )
var NO_SUCH_APP = goErrors.New( "no such app" )
var DUPLICATE_APP = goErrors.New( "duplicate app" )
var SOURCE_ADD_NO_APP = goErrors.New( "source add: no app" )
var SOURCE_DELETE_NO_APP = goErrors.New( "source delete: no app" )
var REPO_INDEX_DOES_NOT_EXIST = goErrors.New( "repo index does not exist" )
var INSTALLED_APPS_INDEX_DOES_NOT_EXIST = goErrors.New( "installed apps index does not exist" )
var INSTALL_DIR_DOES_NOT_EXIST = goErrors.New( "install dir does not exist" )
var APP_SEARCH_NO_SEARCH_TERM = goErrors.New( "app search: no search term" )
var APP_INSTALL_NO_APP_ID = goErrors.New( "app install: no app id" )
var APP_ALREADY_INSTALLED = goErrors.New( "app already installed" )